import axios from '.'
import {Line,Im} from '../interface/index'

export const GetImLine= (im:Im)=>{
    return axios.post<Line[]>("/api/im/line",im)
}