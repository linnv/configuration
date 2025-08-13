#!/usr/bin/env bash
# run-jina-v3-tornado.sh
set -e

IMG_NAME="jina-v3-tornado"
CONT_NAME="jina-v3-tornado-ctr"
HOST_PORT="${HOST_PORT:-8091}"
GPU_FLAG=""

# ---------- 可选参数 ----------
# export https_proxy=http://user:pass@proxy:port
# export USE_GPU=1
# export HOST_PORT=9000
# ------------------------------

# GPU 支持
if [[ "$USE_GPU" == "1" ]]; then
  GPU_FLAG="--gpus all"
  BASE_IMAGE="nvidia/cuda:12.2-devel-ubuntu22.04"
  TORCH_INDEX="https://download.pytorch.org/whl/cu121"
else
  BASE_IMAGE="python:3.12-slim"
  TORCH_INDEX="https://download.pytorch.org/whl/cpu"
fi

# 生成 requirements.txt
cat > requirements.txt <<EOF
torch --index-url $TORCH_INDEX
transformers
einops
sentence-transformers
tornado
numpy
EOF

# 生成 server.py
cat > server.py <<'PY'
import json
import numpy as np
import tornado.ioloop
import tornado.web
from sentence_transformers import SentenceTransformer
import logging

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

print("Loading model...")
model = SentenceTransformer("jinaai/jina-embeddings-v3", trust_remote_code=True)
print("Model loaded.")

def cosine_similarity(a, b):
    a = np.array(a)
    b = np.array(b)
    return float(np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b)))

class EmbedHandler(tornado.web.RequestHandler):
    async def post(self):
        try:
            data = json.loads(self.request.body)

            sentences = data.get("sentences")
            if not sentences or not isinstance(sentences, list) or not all(isinstance(s, str) for s in sentences):
                self.set_status(400)
                self.write({"error": "Missing 'sentences' field or not a list of strings"})
                return

            # Compute embeddings
            embeddings = model.encode(sentences).tolist()

            response = {"embeddings": embeddings}

            # If similarity requested
            if data.get("compute_similarity", False) and len(sentences) == 2:
                sim = cosine_similarity(embeddings[0], embeddings[1])
                response["similarity"] = sim

            self.write(response)

        except json.JSONDecodeError:
            self.set_status(400)
            self.write({"error": "Invalid JSON format"})
        except Exception as e:
            logger.error(f"Error processing request: {e}")
            self.set_status(500)
            self.write({"error": "Internal server error"})

def make_app():
    return tornado.web.Application([
        (r"/embed", EmbedHandler),
    ])

if __name__ == "__main__":
    app = make_app()
    app.listen(8080)
    logger.info("Tornado server started on port 8080")
    tornado.ioloop.IOLoop.current().start()
PY

# 生成 Dockerfile
cat > Dockerfile <<DOCKER
FROM $BASE_IMAGE
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY server.py .
ENV TRANSFORMERS_CACHE=/tmp/transformers
EXPOSE 8080
CMD ["python", "server.py"]
DOCKER

# 构建
echo ">>> Building image $IMG_NAME ..."
docker build -t "$IMG_NAME" .

# 若已存在同名容器，先删除
docker rm -f "$CONT_NAME" 2>/dev/null || true

# 运行
echo ">>> Running on http://localhost:$HOST_PORT"
docker run -d --rm --name "$CONT_NAME" \
  -p "$HOST_PORT":8080 \
  -v ./data:/data \
  -e TRANSFORMERS_CACHE=/data/models \
  ${GPU_FLAG} \
  ${https_proxy:+-e https_proxy=$https_proxy} \
  ${http_proxy:+-e http_proxy=$http_proxy} \
  "$IMG_NAME"

echo ">>> Container started: $CONT_NAME"


