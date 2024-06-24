import express from 'express';
import cors from 'cors';
import connectToMongo from './db.js';
import UserRoutes from './routes/UserRoutes.js'; 
import http from 'http';
import { Server as SocketIO } from 'socket.io';
import dotenv from 'dotenv';
dotenv.config();


const app = express();
app.use(cors());
app.use(express.json());

// Available Routes
app.use("/api/auth", UserRoutes);


const server = http.createServer(app);

connectToMongo();

server.listen(process.env.PORT || 8000, () => {
    console.log(`app listening on Port ${process.env.PORT || 8000}`);
});



const io = new SocketIO(server, {
  cors: {
    origin: "http://localhost:3000",
    credentials: true,
  }
});


io.on("connection", (socket) => {
  console.log(`User ${socket.handshake.query.username} connected.`);

  const username = socket.handshake.query.username;
  socket.join(username);

  socket.on("verify_key_true", (data) => {
    io.to(data.username).emit("key_verified", { username: data.username, status: true });
  });

  socket.on("send_live_location", (data) => {
    io.to(data.username).emit("live_location", { liveLocStatus: data.liveLocStatus, username: data.username });
  });

  socket.on("disconnect", () => {
    console.log(`User ${socket.handshake.query.username} disconnected.`);
  });
});

