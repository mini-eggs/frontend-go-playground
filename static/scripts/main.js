function CreateComponent(props) {
  return class extends React.Component {
    componentDidMount() {
      (props.ComponentDidMount ? props.ComponentDidMount : () => {})();
    }
    render() {
      return (props.Render ? props.Render : () => null)();
    }
  };
}
