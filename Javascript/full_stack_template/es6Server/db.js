import mongoose from 'mongoose';



mongoose.set('strictQuery', false);

const connectToMongo = () => {
  mongoose.connect(process.env.MONGO_URI)
  .then(() => console.log("app connected to MongoDB Successfully"))
  .catch((err) => console.log(err.message))
};

export default connectToMongo;
