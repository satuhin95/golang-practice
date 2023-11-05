import React from "react";
import {  Container} from "react-bootstrap";
import { Link } from "react-router-dom";

const Header = () => {
  return (
    <>
      <Container fluid className="container-fluid header">
      <Link to="/">
        <h3 className="text-center text-uppercase ">
        React Application with Go fiber Backend
        </h3>
        </Link>
      </Container>
      <Container>
        <div>
          {/* <ul className="menu">
            <li>
              <Link to="/">Home</Link>
            </li>
      
          </ul> */}
        </div>
      </Container>
    </>
  );
};

export default Header;