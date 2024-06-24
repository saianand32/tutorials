import bcrypt from 'bcrypt'
import jwt from 'jsonwebtoken'
import User from '../models/UserModel.js'


const signup = async (req, res, next) => {
  try {
    const { username, email, password } = req.body;

    const usernameCheck = await User.findOne({ username });

    if (usernameCheck)
      return res.status(400).json({ msg: "Username already used", status: false });

    const emailCheck = await User.findOne({ email });
    if (emailCheck)
      return res.status(400).json({ msg: "Email already used", status: false });

    const securePassword = await bcrypt.hash(password, 10); // 10 is the salt

    const user = await User.create({
      email,
      username,
      password: securePassword,
    });

    const data = {
      user: {
        id: user._id,
        username: user.username,
        email: user.email,
      },
    };
    const authToken = jwt.sign(data, process.env.JWT_SECRET);

    return res.status(201).json({ status: true, authToken, username });
  } catch (ex) {
    next(ex);
  }
};

const login = async (req, res, next) => {
  try {
    const { username, password } = req.body;
    let user = await User.findOne({ username });
    if (!user) {
      return res.status(401).json({
        status: false,
        msg: "Please login with correct details",
      });
    }
    const passwordCompare = await bcrypt.compare(password, user.password);
    if (!passwordCompare) {
      return res.status(401).json({
        status: false,
        msg: "Please login with correct details",
      });
    }

    const data = {
      user: {
        id: user._id,
        username: user.username,
        email: user.email,
      },
    };
    const authToken = jwt.sign(data, process.env.JWT_SECRET);
    res.json({ status: true, authToken, username });
  } catch (error) {
    console.error(error.message);
    res.status(500).json({ status: false, msg: "Internal server error occurred" });
  }
};

const temp = async (req, res, next) => {
  // just to demo a controller using fetchuser middleware
  try {
    res.json({
      username: req.user.username,
      email: req.user.email,
      id: req.user.id,
    });
  } catch (ex) {
    next(ex);
  }
};

export { signup, login, temp }
