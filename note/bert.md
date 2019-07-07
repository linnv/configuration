[
rm -rf /tmp/qn_output/*
rm -rf /tmp/qn_pout/*
export BERT_BASE_DIR=data/chinese_L-12_H-768_A-12
export TRAINED_OUTPUT=/tmp/qn_output/
export PREDIC_OUTPUT=/tmp/qn_pout/
export GLUE_DIR=/data/py/bert/glue

python3 run_classifier.py \
  --task_name=qn \
  --do_train=true \
  --do_eval=true \
  --data_dir=$GLUE_DIR \
  --vocab_file=$BERT_BASE_DIR/vocab.txt \
  --bert_config_file=$BERT_BASE_DIR/bert_config.json \
  --init_checkpoint=$BERT_BASE_DIR/bert_model.ckpt \
  --max_seq_length=128 \
  --train_batch_size=32 \
  --learning_rate=2e-5 \
  --num_train_epochs=3.0 \
  --output_dir=$TRAINED_OUTPUT

python3 run_classifier.py \
  --task_name=qn \
  --do_predict=true \
  --data_dir=$GLUE_DIR \
  --vocab_file=$BERT_BASE_DIR/vocab.txt \
  --bert_config_file=$BERT_BA
]
