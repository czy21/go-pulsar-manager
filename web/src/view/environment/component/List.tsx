import stub from "@/init"
import React from "react";
import react from "react";
import ClusterAdd from "./Add"
import {Filter, intl, OperationRender, Table} from "@c";

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
            label: <intl.FormatMessage id={"environment.edit"}/>,
            onClick: (text: any, record: any) => {
                console.log(record)
            }
        }
    ]
    const columns = [
        {
            key: 'name',
            header: <intl.FormatMessage id={"environment.name"}/>,
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
            header: <intl.FormatMessage id={"environment.url"}/>,
        },
        {
            key: 'description',
            header: <intl.FormatMessage id={"environment.description"}/>,
        },
        {
            key: 'operation',
            header: <intl.FormatMessage id={"table.operation"}/>,
            render: (text: any, record: any) => OperationRender(text, record, operationActions)
        }
    ];

    const handleSearch = (q?: any) => {
        setQuery(q)
        stub.api.post("environment/search", stub.ref.lodash.omit(q, "total")).then((t: any) => setData({list:t.data.data}))
    }

    const [clusterAddVisible, setClusterAddVisible] = React.useState<boolean>(false);

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
            <stub.ref.antd.Space>
                <stub.ref.antd.Button type={"primary"} onClick={() => setClusterAddVisible(true)}>
                    {<intl.FormatMessage id={"environment.add"}/>}
                </stub.ref.antd.Button>
            </stub.ref.antd.Space>
            <Table columns={columns}
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

export default Index