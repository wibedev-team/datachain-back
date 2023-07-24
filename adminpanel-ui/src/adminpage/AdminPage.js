import React from "react";
import {Route, Routes} from "react-router-dom";
import {Navbar} from "../navbar/navbar";
import {About} from "../about/about";
import {Stack} from "../stack/stack";
import {Solutions} from "../solutions/solutions";
import {Teams} from "../teams/teams";
import {Footer} from "../footer/footer";
import "../navbar/navbar.css"

export const AdminPage = () => {
    return (
        <>
            <Navbar />
            <Routes>
                <Route path={"/about"} element={<About />} />
                <Route path={"/stack"} element={<Stack />} />
                <Route path={"/solutions"} element={<Solutions />} />
                <Route path={"/teams"} element={<Teams />} />
                <Route path={"/footer"} element={<Footer />} />
            </Routes>
        </>
    );
}