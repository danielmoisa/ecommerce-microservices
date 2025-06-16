import React, {useState, useEffect} from "react";
import {Avatar, Button, List, PageHeader, Pagination, Skeleton} from "antd";
import {getProductDetail} from "../../api/catalog";
import {useParams} from "react-router-dom";

export default function ProductDetail() {
    let { id } = useParams();
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
                title="Product Detail"
            />
            <div>
                {id}
            </div>
        </div>
    )
}
