import stub from "@/init";
import React from "react";

import {intl, Modal} from '@c'

interface TableFormProp {
    visible: boolean
    onChange: () => void
}

const Index: React.FC<TableFormProp> = (props: TableFormProp | any) => {
    const [visible, setVisible] = React.useState<boolean>(false)
    const [connectState, setConnectState] = React.useState<boolean>()
    React.useEffect(() => setVisible(props.visible as boolean), [props.visible])

    const [addForm] = stub.ref.antd.Form.useForm();

    const handleOk = () => {
        stub.util.basic.validateForm(addForm.validateFields(),
            (values) => {
                stub.api.post("environment/add", values).then(t => {
                    addForm.resetFields()
                    setConnectState(undefined)
                    props.onChange()
                })
            })
    };

    const handleCancel = () => {
        addForm.resetFields()
        setConnectState(undefined)
        props.onChange()
    }

    const handleConnect = () => {
        stub.util.basic.validateForm(addForm.validateFields(),
            (values) => {
                stub.api.post("environment/ping", values)
                    .then((t: any) => setConnectState(stub.ref.lodash.isEmpty(t.error)))
            })
    }

    return (
        <div>
            <Modal
                title={"添加环境"}
                width={600}
                visible={visible}
                onOk={handleOk}
                onCancel={handleCancel}>
                <stub.ref.antd.Form form={addForm}
                                    labelCol={{span: 8}}
                                    wrapperCol={{span: 10}}
                >
                    <stub.ref.antd.Form.Item label={<intl.FormatMessage id={"environment.name"}/>} name={"name"}
                                             rules={[{required: true}]}
                    >
                        <stub.ref.antd.Input/>
                    </stub.ref.antd.Form.Item>
                    <stub.ref.antd.Form.Item label={<intl.FormatMessage id={"environment.url"}/>} name={"url"}
                                             rules={[{required: true}]}
                    >
                        <stub.ref.antd.Input/>
                    </stub.ref.antd.Form.Item>
                    <stub.ref.antd.Form.Item {...{wrapperCol: {offset: 8}}}>
                        <stub.ref.antd.Space>
                            <stub.ref.antd.Button onClick={handleConnect} size={"small"}>Test</stub.ref.antd.Button>
                            {connectState !== undefined ?
                                (connectState ? <stub.ref.icon.ai.CheckCircleTwoTone twoToneColor={"#52c41a"}/> : <stub.ref.icon.ai.CloseCircleTwoTone twoToneColor={"#ff4d4f"}/>) : ""}
                        </stub.ref.antd.Space>
                    </stub.ref.antd.Form.Item>
                    <stub.ref.antd.Form.Item name={"description"} label={<intl.FormatMessage id={"environment.description"}/>}>
                        <stub.ref.antd.Input.TextArea/>
                    </stub.ref.antd.Form.Item>
                </stub.ref.antd.Form>
            </Modal>
        </div>
    )
}

export default stub.ref.reactRedux.connect(stub.util.basic.mapGlobalOptionStateToProps)(Index)