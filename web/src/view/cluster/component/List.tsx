import stub from "@/init"
import React from "react";
import react from "react";
import Add from "./Add"
import {intl, OperationRender, Table} from "@c";

const Index: React.FC<any> = (props: any) => {
    const homeState = stub.ref.reactRedux.useSelector((state: any) => state.home)
    const [data, setData] = React.useState<any>({})
    const [query, setQuery] = React.useState<any>({})

    react.useEffect(() => {
        handleSearch(query)
    }, [homeState.environment])
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
            primaryKey: true,
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
            key: 'brokers',
            header: <intl.FormatMessage id={"cluster.brokers"}/>,
        },
        {
            key: 'operation',
            header: <intl.FormatMessage id={"table.operation"}/>,
            render: (text: any, record: any) => OperationRender(text, record, operationActions)
        }
    ];

    const handleSearch = (q?: any) => {
        if (homeState.environment.extra) {
            q.serviceUrl = homeState.environment.extra.serviceUrl
            setQuery(q)
            stub.api.post("cluster/search", q).then((t: any) => setData({"list": t.data.data}))
        }
    }

    const [addVisible, setAddVisible] = React.useState<boolean>(false);

    const extension = (
        <stub.ref.antd.Space>
            <stub.ref.antd.Button type={"primary"} onClick={() => setAddVisible(true)}>
                {<intl.FormatMessage id={"cluster.add"}/>}
            </stub.ref.antd.Button>
        </stub.ref.antd.Space>
    )
    const [form] = stub.ref.antd.Form.useForm();
    return (
        <div>
            <stub.ref.antd.Form
                form={form}
                layout="inline"
            >
                <stub.ref.antd.Form.Item>
                    <stub.ref.antd.Button type={"primary"} onClick={() => {
                        stub.util.basic.validateForm(form.validateFields(),
                            (value: any) => {
                                value.serviceUrl = homeState.environment.extra.serviceUrl
                                handleSearch(value)
                            })
                    }}>{<intl.FormatMessage id={"common.filter.search"}/>}
                    </stub.ref.antd.Button>
                </stub.ref.antd.Form.Item>
            </stub.ref.antd.Form>
            <Table columns={columns} list={data.list} page={data.page}/>
            <Add visible={addVisible} onChange={() => {
                setAddVisible(false)
                handleSearch(query)
            }}/>
        </div>
    )

}

export default Index