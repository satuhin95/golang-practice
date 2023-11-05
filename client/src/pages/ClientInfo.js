import axios from 'axios'
import moment from 'moment'
import React, { useEffect, useState } from 'react'
import { Container,Card,ListGroup,Row,Col } from 'react-bootstrap'
import { useParams } from 'react-router-dom'
import Spinner from 'react-bootstrap/Spinner';
function ClientInfo() {
    const [loading, setLoading] = useState(true)
    const [clientInfo, setClientInfo] = useState(false)
    const {id} = useParams()

    useEffect(()=>{
        const fetchData = async ()=>{
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT + "/clientinfo/" + id;
                const response = await axios.get(apiUrl);
  
                if (response.status === 200) {
                    setClientInfo(response?.data?.data);       
                 }
                 setLoading(false)
            } catch (error) {
                setLoading(false)
                console.log(error.response);
            }
        }
        fetchData()
    },[])
    
  if (loading) {
    return (
      <>
        <Container className="spinner">
          <Spinner animation="grow" />
        </Container>
      </>
    );
  }
  return (
    <Container>
      <Row className="justify-content-md-center">
      <Col  className="py-5" md="auto">
    {clientInfo && (
        <Card style={{ width: '30rem' }} classnName="text-center">
             <ListGroup variant="flush">
            <ListGroup.Item>SourceName : {clientInfo.source_name}</ListGroup.Item>
            <ListGroup.Item>Unit : {clientInfo.unit}</ListGroup.Item>
            <ListGroup.Item>Quantity : {clientInfo.quantity}</ListGroup.Item>
            <ListGroup.Item>Count : {clientInfo.count}</ListGroup.Item>
            <ListGroup.Item>Value : {clientInfo.value}</ListGroup.Item>
            <ListGroup.Item>Operating System Version : {clientInfo.operating_system_version}</ListGroup.Item>
            <ListGroup.Item>Source Bundle Identifier : {clientInfo.source_bundle_identifier}</ListGroup.Item>
            <ListGroup.Item>SourceProductType : {clientInfo.source_product_type}</ListGroup.Item>
            <ListGroup.Item>Quantity Type : {clientInfo.quantity_type}</ListGroup.Item>
            <ListGroup.Item>QuantityTypeOther : {clientInfo.quantity_type_other}</ListGroup.Item>
            <ListGroup.Item>StartDate : {moment(clientInfo.start_date).format("llll")}</ListGroup.Item>
            <ListGroup.Item>EndDate : {moment(clientInfo.end_date).format("llll")}</ListGroup.Item>

          </ListGroup>
        </Card>
    )}
      </Col>
    </Row>
    </Container>
  )
}

export default ClientInfo
