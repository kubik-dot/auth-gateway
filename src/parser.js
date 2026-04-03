import { Router } from 'express';
import { validate } from 'schema';
import { sendResponse } from './utils';
import { UserSchema } from './schemas';

const router = Router();

router.post('/login', (req, res) => {
  const { error } = validate(req.body, UserSchema);
  if (error) {
    return sendResponse(res, 400, error.details[0].message);
  }

  const { username, password } = req.body;

  // Simulate database query
  const user = { username, password: 'hashed_password' };

  if (user.username === username && user.password === password) {
    sendResponse(res, 200, { token: 'example_token' });
  } else {
    sendResponse(res, 401, 'Invalid credentials');
  }
});

export default router;