import stub from "@/init"
import React from "react";
import react from 'react'
import NamespaceAdd from "./Add"
import {Table, Filter, OperationRender, intl} from "@c";

const Index: React.FC<any> = (props: any) => {

    const [data, setData] = React.useState<any>({})
    const [query, setQuery] = React.useState<any>()

    react.useEffect(() => {
        // stub.store.dispatch(stub.reducer.action.option.fetch(["dbInstanceKind", "genderKind"]))
        handleSearch()
    }, [])
    const operationActions = [
        {
            key: "edit",
            label: <intl.FormatMessage id={"cluster.edit"}/>,
            onClick: (text: any, record: any) => {
                console.log(record)
            }
        }
    ]
    const columns = [
        {
            key: 'name',
            header: <intl.FormatMessage id={"cluster.name"}/>,
            render: (text: any, record: any) => {
                return (
                    <a onClick={() => {
                        props.history.push(`${props.route.path}/${record.name}`)
                    }}>{record.name}</a>
                )
            }
        },
        {
            key: 'url',
            header: <intl.FormatMessage id={"cluster.url"}/>,
        },
        {
            key: 'description',
            header: <intl.FormatMessage id={"cluster.description"}/>,
        },
        {
            key: 'operation',
            header: <intl.FormatMessage id={"table.operation"}/>,
            render: (text: any, record: any) => OperationRender(text, record, operationActions)
        }
    ];

    const handleSearch = (q?: any) => {
        setQuery(q)
        stub.api.post("cluster/search", stub.ref.lodash.omit(q, "total")).then((t: any) => setData(t.data.data))
    }

    const [clusterAddVisible, setClusterAddVisible] = React.useState<boolean>(false);

    const extension = (
        <stub.ref.antd.Space>
            <stub.ref.antd.Button type={"primary"} onClick={() => setClusterAddVisible(true)}>
                {<intl.FormatMessage id={"cluster.add"}/>}
            </stub.ref.antd.Button>
        </stub.ref.antd.Space>
    )

    return (
        <div>
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
            <Table columns={columns}
                   list={data.list}
                   page={data.page}
            />
            <NamespaceAdd visible={clusterAddVisible} onChange={() => {
                setClusterAddVisible(false)
                handleSearch(query)
            }}/>
        </div>
    )

}

export default Index