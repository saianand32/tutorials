import { Router } from 'express';
import { signup, temp, login } from '../controllers/authController.js';
import fetchuser from './middleware/FetchUser.js';

const router = Router();

router.post("/signup", signup);
router.post("/login", login);
router.get("/temp", fetchuser, temp);

export default router;
