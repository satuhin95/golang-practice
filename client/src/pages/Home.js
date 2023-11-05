import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Container, Row, Spinner,Col } from 'react-bootstrap'
import ClientInfoCard from '../components/ClientInfoCard'

function Home() {
    const [clientInfo, setClientInfo] = useState(false)
    const [loading, setLoading] = useState(true)

    useEffect(()=>{
        const fetchData = async()=>{
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT + "/clientinfo"
                const response = await axios.get(apiUrl)

                if (response.status === 200){
                    setClientInfo(response?.data?.data)
                    console.log(response?.data?.data)
                }
                setLoading(false)
            } catch (error) {
                setLoading(false)
                console.log(error.response)
            }
        }
        fetchData()
        return () => {};
    },[])
    if(loading){
        return(
            <Container className='spinner'>
              <Spinner animation="border" variant="secondary"  size="sm" />
              <Spinner animation="border" variant="success"  size="sm" />
              <Spinner animation="border" variant="danger" size="sm" />
              <Spinner animation="border" variant="warning" size="sm" />
              <Spinner animation="border" variant="info" size="sm" />

            </Container>
          )
    }
  return (
    <Container>
      <Row className="justify-content-md-center">
        {clientInfo && (
            clientInfo.map((client,index)=>(

                <Col  className="py-5" key={index}>
                    <ClientInfoCard  client={client} index={index}/>
                </Col>
            ))
        )}
      </Row>
    </Container>
  )
}

export default Home
