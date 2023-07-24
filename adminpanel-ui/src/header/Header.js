import React from "react";
import "./header.css"
import {useDispatch, useSelector} from "react-redux";
import {userActions} from "../redux/loginUserSlice";

export const Header = () => {
    const login = useSelector((state) => state.user.login)
    const isAuth = useSelector((state) => state.user.isAuth)
    const dispatch = useDispatch()

    const logout = () => {
        localStorage.removeItem("token")
        localStorage.removeItem("login")
        dispatch(userActions.logout())
    }

    return (
        <nav className={"header"}>
            <h1>admin panel, hello {login}</h1>
            {isAuth?
                <button className={"logout-btn"} onClick={logout} type="submit">logout</button>
                :
                <></>
            }
        </nav>
    );
}