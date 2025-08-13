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
