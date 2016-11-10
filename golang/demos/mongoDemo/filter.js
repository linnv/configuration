use ts
db.accounts.save({name: "A", balance: 1000, pendingTransactions: []})
db.accounts.save({name: "B", balance: 1000, pendingTransactions: []})
db.accounts.find()
db.transactions.save({source: "A", destination: "B", value: 100, state: "initial"})
t = db.transactions.findOne({state: "initial"})
db.transactions.update({_id: t._id}, {$set: {state: "pending"}})
db.transactions.find()
db.accounts.update(
{ name: t.source, pendingTransactions: { $ne: t._id } },
{ $inc: { balance: -t.value }, $push: { pendingTransactions: t._id } }
)
db.accounts.update(
{ name: t.destination, pendingTransactions: { $ne: t._id } },
{ $inc: { balance: t.value }, $push: { pendingTransactions: t._id } }
)
db.accounts.find()
db.transactions.update({_id: t._id}, {$set: {state: "committed"}})
db.transactions.find()
db.accounts.find()
db.accounts.update({name: t.source}, {$pull: {pendingTransactions: t._id}})
db.accounts.update({name: t.destination}, {$pull: {pendingTransactions: t._id}})
db.accounts.find()
db.transactions.update({_id: t._id}, {$set: {state: "done"}})
db.transactions.find()

db.transactions.update({_id: t._id}, {$set: {state: "canceling"}})
db.accounts.update({name: t.source, pendingTransactions: t._id}, {$inc: {balance: t.value}, $pull: {pendingTransactions: t._id}})
