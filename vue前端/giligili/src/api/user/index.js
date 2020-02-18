import axios from 'axios';

//用户注册
const register = form => axios.post('/api/v1/user/register', form).then(res => res.data);

//用户登录
const login = form => axios.post('/api/v1/user/login', form).then(res => res.data);

//预定车票
const order = form => axios.post('/api/v1/user/order', form).then(res => res.data);

//输入车票信息
const input = form => axios.post('/api/v1/user/input', form).then(res => res.data);

const search = form => axios.post('/api/v1/user/search', form).then(res => res.data);

export {
    register,
    login,
    order,
    input,
    search
};
