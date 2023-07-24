import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    login: "",
    token: "",
    isAuth: false
}

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setLogin: (state, action) => {
            state.login = action.payload
        },
        setToken: (state, action) => {
            state.token = action.payload
        },
        setIsAuth: (state) => {
            state.isAuth = true
        },
        logout: (state) => {
            state.isAuth = false
        },
    },
})

export const {actions: userActions} = userSlice;
export const {reducer: userReducer} = userSlice;