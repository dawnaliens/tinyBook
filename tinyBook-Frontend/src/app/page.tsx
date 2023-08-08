'use client';

import React from 'react';
import SignupForm from '../pages/users/signup';
import LoginForm from '../pages/users/login';
import { BrowserRouter, Route, Routes } from "react-router-dom"
import {Row} from "antd";

// @ts-ignore
const Signup = () => (
    <div>
        <Row type="flex" justify="center" align={"middle"}>
            <SignupForm></SignupForm>
        </Row>
    </div>
);

const Login = () => (
    <div>
        <Row type="flex" justify="center" align={"middle"}>
            <LoginForm></LoginForm>
        </Row>
    </div>
);



const App = () => {
    return <BrowserRouter>
        <Routes>
            <Route index element={<Signup />} />
            <Route path="signup" element={<Signup />} />
            <Route path="login" element={<Login />} />
        </Routes>
    </BrowserRouter>
}

export default App;