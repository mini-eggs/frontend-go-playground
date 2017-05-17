import "./styles/main.scss";
import "./styles/header.scss";
import { createElement } from "react";
import { render } from "react-dom";
import { Switch, Route } from "react-router";
import CreateComponent from "./scripts/createComponent";

window.React = { createElement: createElement };
window.ReactDOM = { render: render };
window.ReactRouter = { Switch: Switch, Route: Route };
window.ReactRouterDOM = require("react-router-dom");
window.CreateComponent = CreateComponent;
