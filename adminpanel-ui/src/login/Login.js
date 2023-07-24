import React, {useState} from "react";
import "./login.css"
import {useNavigate} from "react-router-dom";
import {useDispatch} from "react-redux";
import {userActions} from "../redux/loginUserSlice";
import instance from "../axios/axios";

export const Login = () => {
    const [login, SetLogin] = useState("")
    const [password, SetPassword] = useState("")
    const navigate = useNavigate()
    const dispatch = useDispatch()

    const handleLogin = () => {
        instance.post("/auth/login",{
            "login": login,
            "password": password,
        }).then(response => {
            console.log(response.data)

            // localStorage.setItem('user', JSON.stringify(response.data.user))
            localStorage.setItem('token', JSON.stringify(response.data.token))
            localStorage.setItem('login', JSON.stringify(response.data.login))

            dispatch(userActions.setLogin(response.data.login))
            dispatch(userActions.setToken(response.data.token))
            dispatch(userActions.setIsAuth())

            navigate('/admin')
        }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <div>
            <h1>Login</h1>
            <div className={"form"}>
                <label className={'auth-label'}>
                    Введите имя логин
                    <input
                        className={'form-control'}
                        type="text"
                        onChange={event => SetLogin(event.target.value)}
                        value={login}
                    />
                </label>
                <label className={'auth-label'}>
                    Введите пароль
                    <input
                        className={'form-control'}
                        type="password"
                        onChange={event => SetPassword(event.target.value)}
                        value={password}
                    />
                </label>
                <button
                    className={'btn'}
                    onClick={handleLogin}
                >
                    ок
                </button>
            </div>
        </div>
    );
}