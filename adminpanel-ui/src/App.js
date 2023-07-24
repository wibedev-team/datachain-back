import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import React, {useEffect} from "react";
import {Login} from "./login/Login";
import {AdminPage} from "./adminpage/AdminPage";
import {Header} from "./header/Header";
import {useDispatch, useSelector} from "react-redux";
import {userActions} from "./redux/loginUserSlice";
import {About} from "./about/about";
import {Footer} from "./footer/footer";
import {Teams} from "./teams/teams";
import {Solutions} from "./solutions/solutions";
import {Stack} from "./stack/stack";

export const App = () => {
    const isAuth = useSelector((state) => state.user.isAuth)
    const dispatch = useDispatch()

    useEffect(() => {
        const token = JSON.parse(localStorage.getItem('token'))
        if (!token) {
            dispatch(userActions.logout)
        } else {
            const login = JSON.parse(localStorage.getItem('login'))
            dispatch(userActions.setLogin(login))
            dispatch(userActions.setToken(token))
            dispatch(userActions.setIsAuth())
        }
    }, [dispatch]);

  return (
      <BrowserRouter>
        <Header />
        <Routes>
          <Route path={'/'} element={isAuth ? <Navigate to={'/admin'} /> : <Login />} />
          <Route path={'/login'} element={isAuth ? <Navigate to={'/admin'} /> : <Login />} />
          <Route path={'admin'} element={isAuth ? <AdminPage /> : <Navigate to={'/login'} />}>
             <Route index path={'about'} element={<About />} />
             <Route index path={'stack'} element={<Stack />} />
             <Route index path={'solutions'} element={<Solutions />} />
             <Route index path={'teams'} element={<Teams />} />
             <Route index path={'footer'} element={<Footer />} />
          </Route>
        </Routes>
      </BrowserRouter>
  );
}