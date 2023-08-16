import React, {useEffect, useState} from "react";
import "./index.css"
import instance, {MINIO} from "../axios/axios";
import {useNavigate} from "react-router-dom";

export const Stack = () => {
    const [images, SetImages] = useState([])
    const [file, SetFile] = useState(null)
    const [enable, SetEnable] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        instance.get("/stack/all")
            .then(response => {

                console.log(response.data)
                SetImages(response.data.stacks)

            }).catch((error) => {
            console.log(error);
        })
    }, [])

    const handleImage = (e) => {
        SetFile(e.target.files[0])
        SetEnable(true)
    }

    const createImg = () => {
        const formData = new FormData()
        formData.append('image', file)
        instance.post("/stack/create", formData, {
            headers: {Authorization: `Bearer ${JSON.parse(localStorage.getItem("token"))}`
        }}).then(response => {
            console.log(response.data)
            navigate("/admin")
        }).catch((error) => {
            console.log(error);
        })
    }

    const remove = (id) => {
        instance.delete(`/stack/${id}`, {
            headers: {Authorization: `Bearer ${JSON.parse(localStorage.getItem("token"))}`
        }})
            .then(response => {
                console.log(response.data)
                navigate("/admin")

            }).catch((error) => {
            console.log(error);
        })
    }

    return (
       <div>
           <div className={"stack-form"}>
               <input type="file" onChange={e => handleImage(e)}/>
               <button disabled={!enable} onClick={createImg}>добавить</button>
           </div>
           {images !== null && images.map(img =>
               <div key={img.img}>
                   <img className={"stack-img"}  src={MINIO+img.img} alt=""/>
                   <button onClick={() => remove(img.img)}>удалить</button>
               </div>
           )}
       </div>
    );
}