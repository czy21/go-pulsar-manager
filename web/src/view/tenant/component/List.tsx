import stub from "@/init"
import React from "react";
import react from "react";
import {intl, OperationRender, Table} from "@c";
import Add from "./Add";

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
            label: <intl.FormatMessage id={"tenant.edit"}/>,
            onClick: (text: any, record: any) => {
                console.log(record)
            }
        }
    ]
    const columns = [
        {
            key: 'name',
            primaryKey: true,
            header: <intl.FormatMessage id={"tenant.name"}/>,
            render: (text: any, record: any) => {
                return (
                    <a onClick={() => {
                        props.history.push(`${props.route.path}/${record.name}`)
                    }}>{record.name}</a>
                )
            }
        },
        {
            key: 'namespaces',
            header: <intl.FormatMessage id={"tenant.namespaces"}/>,
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
            stub.api.post("tenant/search", q).then((t: any) => setData({"list": t.data.data}))
        }
    }

    const [addVisible, setAddVisible] = React.useState<boolean>(false);

    const extension = (
        <stub.ref.antd.Space>
            <stub.ref.antd.Button type={"primary"} onClick={() => setAddVisible(true)}>
                {<intl.FormatMessage id={"tenant.add"}/>}
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
            <Table columns={columns} list={data.list}/>
            <Add visible={addVisible} onChange={() => {
                setAddVisible(false)
                handleSearch(query)
            }}/>
        </div>
    )

}

export default Index