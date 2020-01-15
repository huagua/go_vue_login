import axios from 'axios';

//用户注册
const register = form => axios.post('/api/v1/user/register', form).then(res => res.data);

//用户登录
const login = form => axios.post('/api/v1/user/login', form).then(res => res.data);

export {
    register,
    login,
};
