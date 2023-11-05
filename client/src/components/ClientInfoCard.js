import Card from 'react-bootstrap/Card';
import moment from 'moment'
import ListGroup from 'react-bootstrap/ListGroup';

import { Link } from 'react-router-dom';

function ClientInfoCard({client,index}) {
  return (
    <Card style={{ width: '20rem' }} key={index}>
    <ListGroup variant="flush">
      <ListGroup.Item>SourceName : {client.source_name}</ListGroup.Item>
      <ListGroup.Item>SourceProductType : {client.source_product_type}</ListGroup.Item>
      <ListGroup.Item>QuantityTypeOther : {client.quantity_type_other}</ListGroup.Item>
      <ListGroup.Item>Value : {client.value}</ListGroup.Item>
      <ListGroup.Item>StartDate : {moment(client.start_date).format("llll")}</ListGroup.Item>
      <ListGroup.Item>EndDate : {moment(client.end_date).format("llll")}</ListGroup.Item>

    </ListGroup>

    <Link className='btn btn-primary' to={`clientinfo/${client.ID}`}> See Details</Link>
  </Card>

  )
}

export default ClientInfoCard
