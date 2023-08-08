import { useState, useEffect } from 'react'
import axios from "@/axios/axios";

function Page() {
    type Profile = {
        Email: string
        Phone: string
    }
    let p: Profile = {Email: "", Phone: ""}
    const [data, setData] = useState<Profile>(p)
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        axios.get('/users/profile')
            .then((res) => res.data)
            .then((data) => {
                setData(data)
                setLoading(false)
            })
    }, [])

    if (isLoading) return <p>Loading...</p>
    if (!data) return <p>No profile data</p>

    return (
        <div>
            <h1>{data.Email}</h1>
            <h1>{data.Phone}</h1>
            {/*<p>{data.bio}</p>*/}
        </div>
    )
}

export default Page