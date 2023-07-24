import React, {useEffect, useState} from "react";
import "./index.css"
import {useNavigate} from "react-router-dom";
import instance, {MINIO} from "../axios/axios";

export const Teams = () => {
    const [teammates, SetTeammates] = useState([])
    const [enable, SetEnable] = useState(false)

    const [name, SetName] = useState("")
    const [position, SetPos] = useState("")
    const [link, SetLink] = useState("")
    const [file, SetFile] = useState(null)

    const navigate = useNavigate()

    useEffect(() => {
        instance.get("/team/get")
            .then(response => {

                console.log(response.data)
                SetTeammates(response.data.teammates)

            }).catch((error) => {
            console.log(error);
        })
    }, [])

    const handleImg = (e) => {
        SetFile(e.target.files[0])
        SetEnable(true)
    }

    const saveTeammate = () => {
        const formData = new FormData()
        formData.append('image', file)
        formData.append('name', name)
        formData.append('position', position)
        formData.append('link', link)
        instance.post("/team/create", formData)
            .then(response => {

                console.log(response.data)
                navigate("/admin")

            }).catch((error) => {
            console.log(error);
        })
    }

    const remove = (id) => {
        instance.delete(`/team/${id}`)
            .then(response => {
                console.log(response.data)
                navigate("/admin")
            }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <div>
            <h1>teams</h1>
            <div className={"team-form"}>
                <label>name</label>
                <input type="text" value={name} onChange={e => SetName(e.target.value)}/>
                <label>position</label>
                <input type="text" value={position} onChange={e => SetPos(e.target.value)}/>
                <label>link</label>
                <input type="text" value={link} onChange={e => SetLink(e.target.value)}/>
                <label>img</label>
                <input type="file" onChange={e => handleImg(e)}/>
                <button disabled={!enable} className={"footer-btn"} onClick={saveTeammate}>ok</button>
            </div>
            {teammates != null && teammates.map(t =>
                    <div key={t.img}>
                        <div>Name: {t.name}</div>
                        <div>Position: {t.position}</div>
                        <div>Link: {t.link}</div>
                        <img src={MINIO+t.img} alt={"img"} />
                        <button onClick={() => remove(t.img)}>удалить</button>
                    </div>
            )}
        </div>
    );
}