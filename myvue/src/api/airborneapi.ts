import axios  from "./"
import {Airborne} from "../interface/index"

export const GetAirborneList = () =>{
    return axios.get<Airborne[]>("/api/airborne")
}

export const AddAirborne = (airborne:Airborne) =>{
    return axios.post("/api/airborne",airborne)
}

export const EditAirborne = (airborne:Airborne) => {
    return axios.put("/api/airborne",airborne)
}

export const DeleteAirborne = (airborne:Airborne) => {
    return axios.delete("/api/airborne",airborne)
}