//const Kafka = require("kafkajs").Kafka
const {Kafka} = require("kafkajs")
run();
async function run(){
   try{
      const kafka = new Kafka({
         "clientId":"myapp",
         "brokers":["manu:2181"]
      })

      const admin = kafka.admin();
      console.log("Connecting....")
      await admin.connect()
      console.log("Connected!")
      //A-M, N-Z
      await admin.createTopics({
         "topics": [{
            "topic":"Users",
            "numPartitions":2
         }]
      })
      console.log("Created Successfully!")
      await admin.disconnect();

   }catch(ex){
      console.error('something bad happend ${ex}')
   }finally{
      process.exit(0)
   }
}