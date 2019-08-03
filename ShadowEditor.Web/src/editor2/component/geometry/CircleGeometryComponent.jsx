import { PropertyGrid, PropertyGroup, TextProperty, DisplayProperty, CheckBoxProperty, NumberProperty, IntegerProperty } from '../../../third_party';
import SetGeometryCommand from '../../../command/SetGeometryCommand';

/**
 * 圆形组件
 * @author tengge / https://github.com/tengge1
 */
class CircleGeometryComponent extends React.Component {
    constructor(props) {
        super(props);

        this.selected = null;

        this.state = {
            show: false,
            expanded: true,
            radius: 1.0,
            segments: 16,
            thetaStart: 0.0,
            thetaLength: Math.PI * 2,
        };

        this.handleExpand = this.handleExpand.bind(this);
        this.handleUpdate = this.handleUpdate.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    render() {
        const { show, expanded, radius, segments, thetaStart, thetaLength } = this.state;

        if (!show) {
            return null;
        }

        return <PropertyGroup title={L_GEOMETRY_COMPONENT} show={show} expanded={expanded} onExpand={this.handleExpand}>
            <NumberProperty name={'radius'} label={L_RADIUS} value={radius} onChange={this.handleChange}></NumberProperty>
            <IntegerProperty name={'segments'} label={L_SEGMENTS} value={segments} onChange={this.handleChange}></IntegerProperty>
            <NumberProperty name={'thetaStart'} label={L_THETA_START} value={thetaStart} onChange={this.handleChange}></NumberProperty>
            <NumberProperty name={'thetaLength'} label={L_THETA_LENGTH} value={thetaLength} onChange={this.handleChange}></NumberProperty>
        </PropertyGroup>;
    }

    componentDidMount() {
        app.on(`objectSelected.CircleGeometryComponent`, this.handleUpdate);
        app.on(`objectChanged.CircleGeometryComponent`, this.handleUpdate);
    }

    handleExpand(expanded) {
        this.setState({
            expanded,
        });
    }

    handleUpdate() {
        const editor = app.editor;

        if (!editor.selected || !(editor.selected instanceof THREE.Mesh) || !(editor.selected.geometry instanceof THREE.CircleBufferGeometry)) {
            this.setState({
                show: false,
            });
            return;
        }

        this.selected = editor.selected;

        const { radius, segments, thetaStart, thetaLength } = Object.assign({},
            this.selected.geometry.parameters, {
                show: true,
            });

        this.setState({
            show: true,
            radius: radius === undefined ? 1 : radius,
            segments: segments === undefined ? 8 : segments,
            thetaStart: thetaStart === undefined ? 0 : thetaStart,
            thetaLength: thetaLength === undefined ? Math.PI * 2 : thetaLength,
        });
    }

    handleChange(value, name) {
        this.setState({
            [name]: value,
        });

        if (value === null) {
            return;
        }

        const { radius, segments, thetaStart, thetaLength } = Object.assign({}, this.state, {
            [name]: value,
        });

        app.editor.execute(new SetGeometryCommand(this.selected, new THREE.CircleBufferGeometry(
            radius,
            segments,
            thetaStart,
            thetaLength,
        )));

        app.call(`objectChanged`, this, this.selected);
    }
}

export default CircleGeometryComponent;