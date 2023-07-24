import React from "react";
import {Link} from "react-router-dom";
import "./navbar.css"

export const Navbar = () => {
    return (
        <nav className={"routes"}>
            <Link to={"/admin/about"}>about</Link>
            <Link to={"/admin/stack"}>stack</Link>
            <Link to={"/admin/solutions"}>solutions</Link>
            <Link to={"/admin/teams"}>team</Link>
            <Link to={"/admin/footer"}>footer</Link>
        </nav>
    );
}