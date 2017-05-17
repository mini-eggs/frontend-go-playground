import { Component } from "react";

export default function CreateComponent(props) {
  return class extends Component {
    componentDidMount() {
      (props.ComponentDidMount ? props.ComponentDidMount : () => {})();
    }
    render() {
      return (props.Render ? props.Render : () => null)();
    }
  };
}
