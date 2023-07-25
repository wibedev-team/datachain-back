import React, {useEffect, useState} from "react";
import "./index.css"
import axios from "axios";
import {useNavigate} from "react-router-dom";
import instance from "../axios/axios";

export const Footer = () => {
    const [email, SetEmail] = useState("")
    const [telephone, SetTelephone] = useState("")
    const [address, SetAddress] = useState("")

    const [fromServerEmail, SetFromServerEmail] = useState("")
    const [fromServerTelephone, SetFromServerTelephone] = useState("")
    const [fromServerAddress, SetFromServerAddress] = useState("")

    const [updateEmail, SetUpdateEmail] = useState("")
    const [updateTelephone, SetUpdateTelephone] = useState("")
    const [updateAddress, SetUpdateAddress] = useState("");

    const [edit, setEdit] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        instance.get("/footer/get")
            .then(response => {
                console.log(response.data.footer)

                SetFromServerEmail(response.data.footer.email)
                SetFromServerTelephone(response.data.footer.telephone)
                SetFromServerAddress(response.data.footer.address)

                SetUpdateEmail(response.data.footer.email)
                SetUpdateTelephone(response.data.footer.telephone)
                SetUpdateAddress(response.data.footer.address)

            }).catch((error) => {
            console.log(error);
        })
    }, [])

    const saveFooter = () => {
        instance.post("/footer/create",{
            "email": email,
            "telephone": telephone,
            "address": address,
        }).then(response => {
            console.log(response.data)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    const updateFooter = () => {
        instance.post("/footer/create",{
            "email": updateEmail,
            "telephone": updateTelephone,
            "address": updateAddress,
        }).then(response => {
            console.log(response.data)
            setEdit(false)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }
    
    return (
        <div>
            <h1>footer</h1>
            {fromServerEmail && fromServerAddress && fromServerTelephone ?
                <div className={"footer-form"}>
                    {edit ?
                        <div className={"footer-form"}>
                            <label>email</label>
                            <input type="text" value={updateEmail} onChange={e => SetUpdateEmail(e.target.value)}/>
                            <label>telephone</label>
                            <input type="tel" value={updateTelephone} onChange={e => SetUpdateTelephone(e.target.value)}/>
                            <label>address</label>
                            <input type="text" value={updateAddress} onChange={e => SetUpdateAddress(e.target.value)}/>
                            <div>
                                <button style={{marginRight: "15px"}} className={"footer-btn"} onClick={updateFooter}>сохранить</button>
                                <button className={"footer-btn"} onClick={() => setEdit(false)}>отменить</button>
                            </div>
                        </div>
                    :
                        <>
                            <div>Email: {fromServerEmail}</div>
                            <div>Telephone: {fromServerTelephone}</div>
                            <div>Address: {fromServerAddress}</div>
                            <button className={"footer-btn"} onClick={() => setEdit(true)}>изменить</button>
                        </>

                    }
                </div>
            :
                <div className={"footer-form"}>
                    <label>email</label>
                    <input type="text" value={email} onChange={e => SetEmail(e.target.value)}/>
                    <label>telephone</label>
                    <input type="text" value={telephone} onChange={e => SetTelephone(e.target.value)}/>
                    <label>address</label>
                    <input type="text" value={address} onChange={e => SetAddress(e.target.value)}/>
                    <button className={"footer-btn"} onClick={saveFooter}>ok</button>
                </div>
            }

        </div>
    );
}