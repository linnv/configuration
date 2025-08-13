# L
curl https://api.jina.ai/v1/embeddings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer jina_61508dd3543e4a3a9e01733f84375c1cR25-XpnQQsoJ4HZX-xewDiSzE0x" \
  -d @- <<EOFEOF
  {
    "model": "jina-embeddings-v3",
    "task": "text-matching",
    "input": [
           "如何申请退款？",
    "退款需要多长时间？",
    "退款流程是什么？",
    "如何联系客服？"
    ]
  }
EOFEOF
