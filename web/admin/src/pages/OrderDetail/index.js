import React, {useEffect, useState} from "react";
import {PageHeader} from "antd";
import {useParams} from "react-router-dom";
import {getProductDetail} from "../../api/catalog";

export default function OrderDetail() {
    let {id} = useParams();
    const [data, setData] = useState({});

    useEffect(() => {
        getProductDetail(id).then((res) => {
            setData(res.data)
        });

    }, []);

    return (
        <div>
            <PageHeader
                ghost={false}
                onBack={() => window.history.back()}
                title="Order Detail"
            />
            <div>
                {data.id}
            </div>
        </div>
    )
}
