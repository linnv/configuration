> use ts
switched to db ts
> db.accounts.save({name: "A", balance: 1000, pendingTransactions: []})
WriteResult({ "nInserted" : 1 })
> db.accounts.save({name: "B", balance: 1000, pendingTransactions: []})
WriteResult({ "nInserted" : 1 })
> db.accounts.find()
{ "_id" : ObjectId("580d75d9f7a7dab78c3b5d11"), "name" : "A", "balance" : 1000, "pendingTransactions" : [ ] }
{ "_id" : ObjectId("580d75dcf7a7dab78c3b5d12"), "name" : "B", "balance" : 1000, "pendingTransactions" : [ ] }
> db.transactions.save({source: "A", destination: "B", value: 100, state: "initial"})
WriteResult({ "nInserted" : 1 })
> t = db.transactions.findOne({state: "initial"})
{
	"_id" : ObjectId("580d75fff7a7dab78c3b5d13"),
	"source" : "A",
	"destination" : "B",
	"value" : 100,
	"state" : "initial"
}
> db.transactions.update({_id: t._id}, {$set: {state: "pending"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.transactions.find()
{ "_id" : ObjectId("580d75fff7a7dab78c3b5d13"), "source" : "A", "destination" : "B", "value" : 100, "state" : "pending" }
> db.accounts.update(
... { name: t.source, pendingTransactions: { $ne: t._id } },
... { $inc: { balance: -t.value }, $push: { pendingTransactions: t._id } }
... )
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.accounts.update(
... { name: t.destination, pendingTransactions: { $ne: t._id } },
... { $inc: { balance: t.value }, $push: { pendingTransactions: t._id } }
... )
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.accounts.find()
{ "_id" : ObjectId("580d75d9f7a7dab78c3b5d11"), "name" : "A", "balance" : 900, "pendingTransactions" : [ ObjectId("580d75fff7a7dab78c3b5d13") ] }
{ "_id" : ObjectId("580d75dcf7a7dab78c3b5d12"), "name" : "B", "balance" : 1100, "pendingTransactions" : [ ObjectId("580d75fff7a7dab78c3b5d13") ] }
> db.transactions.update({_id: t._id}, {$set: {state: "committed"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.transactions.find()
{ "_id" : ObjectId("580d75fff7a7dab78c3b5d13"), "source" : "A", "destination" : "B", "value" : 100, "state" : "committed" }
> db.accounts.find()
{ "_id" : ObjectId("580d75d9f7a7dab78c3b5d11"), "name" : "A", "balance" : 900, "pendingTransactions" : [ ObjectId("580d75fff7a7dab78c3b5d13") ] }
{ "_id" : ObjectId("580d75dcf7a7dab78c3b5d12"), "name" : "B", "balance" : 1100, "pendingTransactions" : [ ObjectId("580d75fff7a7dab78c3b5d13") ] }
> db.accounts.update({name: t.source}, {$pull: {pendingTransactions: t._id}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.accounts.update({name: t.destination}, {$pull: {pendingTransactions: t._id}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.accounts.find()
{ "_id" : ObjectId("580d75d9f7a7dab78c3b5d11"), "name" : "A", "balance" : 900, "pendingTransactions" : [ ] }
{ "_id" : ObjectId("580d75dcf7a7dab78c3b5d12"), "name" : "B", "balance" : 1100, "pendingTransactions" : [ ] }
> db.transactions.update({_id: t._id}, {$set: {state: "done"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.transactions.find()
{ "_id" : ObjectId("580d75fff7a7dab78c3b5d13"), "source" : "A", "destination" : "B", "value" : 100, "state" : "done" }
> db.transactions.update({_id: t._id}, {$set: {state: "canceling"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.accounts.update({name: t.source, pendingTransactions: t._id}, {$inc: {balance: t.value}, $pull: {pendingTransactions: t._id}})
WriteResult({ "nMatched" : 0, "nUpserted" : 0, "nModified" : 0 })
