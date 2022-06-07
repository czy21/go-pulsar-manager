import stub from "@/init"
import React from "react";
import react from 'react'
import ClusterAdd from "./ClusterAdd"
import {Table, Filter, OperationRender} from "@c";

const ClusterList: React.FC<any> = (props: any) => {

    const [data, setData] = React.useState<any>({})
    const [query, setQuery] = React.useState<any>()

    react.useEffect(() => {
        // stub.store.dispatch(stub.reducer.action.option.fetch(["dbInstanceKind", "genderKind"]))
        handleSearch()
    }, [])
    const operationActions = [
        {
            key: "edit",
            label: <stub.ref.intl.FormattedMessage id={"cluster.edit"} defaultMessage={""}/>,
            onClick: (text: any, record: any) => {
                console.log(record)
            }
        }
    ]
    const columns = [
        {
            key: 'name',
            header: <stub.ref.intl.FormattedMessage id={"cluster.name"} defaultMessage={""}/>,
            render: (text: any, record: any) => {
                return (
                    <a onClick={() => {
                        props.history.push(`${props.route.path}/${record.name}`)
                    }}>{record.name}</a>
                )
            }
        },
        {
            key: 'host',
            header: <stub.ref.intl.FormattedMessage id={"cluster.host"} defaultMessage={""}/>,
        },
        {
            key: 'description',
            header: <stub.ref.intl.FormattedMessage id={"cluster.description"} defaultMessage={""}/>,
        },
        {
            key: 'operation',
            header: <stub.ref.intl.FormattedMessage id={"table.operation"} defaultMessage={""}/>,
            render: (text: any, record: any) => OperationRender(text, record, operationActions)
        }
    ];

    const handleSearch = (q?: any) => {
        setQuery(q)
        setData({
            page:{
                pageIndex:1,
                pageSize:10
            },
            list:[
                {
                    name:"集群1",
                    host:"地址1"
                }
            ]
        })
        // stub.api.post("db/instance/search", stub.ref.lodash.omit(q, "total")).then((t: any) => setData(t.data))
    }

    const [clusterAddVisible, setClusterAddVisible] = React.useState<boolean>(false);

    const filter = (
        <Filter
            filters={[
                {
                    "key": "name",
                    "label": "名称"
                }
            ]}
            onSearch={handleSearch}
            page={data.page}
        />
    )

    const extension = (
        <stub.ref.antd.Space>
            <stub.ref.antd.Button type={"primary"} onClick={() => setClusterAddVisible(true)}>
                {<stub.ref.intl.FormattedMessage id={"cluster.add"} defaultMessage={""}/>}
            </stub.ref.antd.Button>
        </stub.ref.antd.Space>
    )

    return (
        <div>
            <Table filter={filter}
                   extension={extension}
                   columns={columns}
                   list={data.list}
                   page={data.page}
            />
            <ClusterAdd visible={clusterAddVisible} onChange={() => {
                setClusterAddVisible(false)
                handleSearch(query)
            }}/>
        </div>
    )

}

export default ClusterList