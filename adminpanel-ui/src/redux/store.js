import {configureStore} from '@reduxjs/toolkit'
import {userReducer, userSlice} from "./loginUserSlice";

export const store = configureStore({
    reducer: {
        user: userReducer,
    },
})

