import React, {useEffect, useState} from "react";
import instance, {MINIO} from "../axios/axios";
import {useNavigate} from "react-router-dom";
import "./index.css"

export const Solutions = () => {
    const [solutions, SetSolutions] = useState([])
    const [title, SetTitle] = useState("")
    const [features, SetFeatures] = useState("")
    const [link, SetLink] = useState("")
    const [file, SetFile] = useState(null)

    const [edit, setEdit] = useState(false)
    const [enable, setEnable] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        instance.get("/solution/all")
            .then(response => {
                console.log(response.data)
                SetSolutions(response.data)

            }).catch((error) => {
            console.log(error);
        })
    }, [])

    const handleImage = (event) => {
        SetFile(event.target.files[0])
        setEnable(true)
    }

    const saveSolution = () => {
        const formData = new FormData()
        formData.append('file', file)
        formData.append('title', title)
        formData.append('features', features)
        formData.append('link', link)

        instance.post("/solution/create", formData, {
            headers: {Authorization: `Bearer ${JSON.parse(localStorage.getItem("token"))}`
        }}).then(response => {
            console.log(response.data)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    const deleteSolution = (id) => {
        instance.delete(`/solution/${id}`, {
            headers: {Authorization: `Bearer ${JSON.parse(localStorage.getItem("token"))}`
        }}).then(response => {
            console.log(response.data)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <div>
            <div>solutions</div>
            <div className={"solutions-form"}>
                <div>Title: каждый уникальный <input value={title} type="text" onChange={e => SetTitle(e.target.value)}/></div>
                <div>Features: вводить каждую фичу с новой строки<textarea onChange={e => SetFeatures(e.target.value)} value={features} style={{resize: "both"}}/></div>
                <div>Link: <input onChange={e => SetLink(e.target.value)} value={link} type="text"/></div>
                <div>File: <input onChange={e => handleImage(e)} type="file"/></div>
                <button disabled={!enable} onClick={saveSolution}>ok</button>
            </div>
            <div>
                {solutions!== null && solutions.map(s =>
                    <div className={"solution"} key={s.title}>
                        <div>{s.title}</div>
                        <ul>{s.features.map(f =>
                            <li>{f.text}</li>
                        )}</ul>
                        <div>{s.link}</div>
                        <div><img src={MINIO+s.file} alt="img"/></div>
                        <button onClick={() => deleteSolution(s.title)}>delete</button>
                    </div>
                )}
            </div>
        </div>

    );
}