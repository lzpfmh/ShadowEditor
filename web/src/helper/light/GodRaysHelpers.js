import BaseHelper from '../BaseHelper';
import GodRays from '../../postprocessing/GodRays';

/**
 * 神光帮助器
 * @author tengge / https://github.com/tengge1
 */
function GodRaysHelpers() {
    BaseHelper.call(this);
}

GodRaysHelpers.prototype = Object.create(BaseHelper.prototype);
GodRaysHelpers.prototype.constructor = GodRaysHelpers;

GodRaysHelpers.prototype.start = function () {
    this.ready = false;
    this.ray = new GodRays();
    this.ray.init(app.editor.scene, app.editor.camera, app.editor.renderer).then(() => {
        this.ready = true;
    });
    app.on(`afterRender.${this.id}`, this.onAfterRender.bind(this));
};

GodRaysHelpers.prototype.stop = function () {
    this.ready = false;
    this.ray.dispose();
    app.on(`afterRender.${this.id}`, null);
};

GodRaysHelpers.prototype.onAfterRender = function () {
    if (!this.ready) {
        return;
    }
    this.ray.render();
};

export default GodRaysHelpers;