import React, {useEffect, useState} from "react";
import "./index.css"
import {useNavigate} from "react-router-dom";
import instance, {MINIO} from "../axios/axios";

export const About = () => {
    const [title, SetTitle] = useState("")
    const [descr, SetDescr] = useState("")
    const [file, setFile] = useState(null);

    const [serverTitle, SetServerTitle] = useState("")
    const [serverDescr, SetServerDescr] = useState("")
    const [serverFile, setServerFile] = useState("");

    const [updateTitle, SetUpdateTitle] = useState("")
    const [updateDescr, SetUpdateDescr] = useState("")
    const [updateFile, SetUpdateFile] = useState(null);

    const [edit, setEdit] = useState(false)
    const [enable, setEnable] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        instance.get("/about/get")
        .then(response => {
            console.log(response.data)
            SetServerTitle(response.data.title)
            SetServerDescr(response.data.description)
            setServerFile(response.data.img)

            SetUpdateTitle(response.data.title)
            SetUpdateDescr(response.data.description)
            SetUpdateFile(response.data.img)

        }).catch((error) => {
            console.log(error);
        })
    }, [])

    const createAbout = () => {
        const formData = new FormData()
        formData.append('image', file)
        formData.append('title', title)
        formData.append('description', descr)

        instance.post("/about/create", formData).then(response => {
            console.log(response.data)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    const updateAbout = () => {
        const formData = new FormData()
        formData.append('image', updateFile)
        formData.append('title', updateTitle)
        formData.append('description', updateDescr)

        instance.post("/about/create", formData).then(response => {
            console.log(response.data)
            setEdit(false)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    const handleFile = (e) => {
        setFile(e.target.files[0])
        setEnable(true)
    }

    const handleUpdateFile = (e) => {
        SetUpdateFile(e.target.files[0])
        setEnable(true)
    }

    return (
        <div>
            <div>About Us section</div>
            {serverTitle ?

                <div>
                    {!edit ?
                        <>
                        <div className={"result"}>
                            <p>{serverTitle}</p>
                            <div dangerouslySetInnerHTML={{__html: serverDescr}} className={"description"} />
                            <img src={MINIO+serverFile} alt="img"/>
                        </div>
                        </>
                        :
                        <>
                            <div className={"edit-form"}>
                                <input type={"text"} value={updateTitle} onChange={e => SetUpdateTitle(e.target.value)} />
                                <textarea value={updateDescr} onChange={e => SetUpdateDescr(e.target.value)} />
                                <input accept={"image/"} type={"file"} onChange={e => handleUpdateFile(e)} />
                            </div>
                        </>
                    }
                    {edit ? <button disabled={!enable} onClick={updateAbout}>сохранить</button>
                        : <button onClick={() => setEdit(true)}>изменить</button>
                    }
                    {edit ? <button onClick={() => setEdit(false)}>отменить</button> : <></>}
                </div>


                : <div className={"about-form"}>
                <label>
                    Введите заголовок
                    <input
                        type="text"
                        onChange={event => SetTitle(event.target.value)}
                        value={title}
                    />
                </label>
                <label className={'auth-label'}>
                    Введите описание
                    <textarea
                        className={'form-control'}
                        onChange={event => SetDescr(event.target.value)}
                        value={descr}
                        style={{resize: "both"}}
                    />
                </label>
                <label className={'auth-label'}>
                    Выберите картинку
                    <input
                        className={'form-control'}
                        type={"file"}
                        onChange={(e) => handleFile(e)}
                    />
                </label>
                <button
                    className={'btn'}
                    onClick={createAbout}
                    disabled={!enable}
                >
                    ок
                </button>
            </div>}
        </div>
    );
}