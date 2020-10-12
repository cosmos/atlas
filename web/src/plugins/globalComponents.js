import Badge from "../components/Badge";
import BaseAlert from "../components/BaseAlert";
import BaseButton from "../components/BaseButton";
import BaseInput from "../components/BaseInput";
import BaseCheckbox from "../components/BaseCheckbox";
import BaseRadio from "../components/BaseRadio";
import BaseSwitch from "../components/BaseSwitch";
import BaseSlider from "../components/BaseSlider";
import BaseProgress from "../components/BaseProgress";
import BasePagination from "../components/BasePagination";
import Card from "../components/Card";

import lang from "element-ui/lib/locale/lang/en";
import locale from "element-ui/lib/locale";

locale.use(lang);

export default {
  install(Vue) {
    Vue.component(Badge.name, Badge);
    Vue.component(BaseAlert.name, BaseAlert);
    Vue.component(BaseButton.name, BaseButton);
    Vue.component(BaseInput.name, BaseInput);
    Vue.component(BaseCheckbox.name, BaseCheckbox);
    Vue.component(BaseRadio.name, BaseRadio);
    Vue.component(BaseSwitch.name, BaseSwitch);
    Vue.component(BaseSlider.name, BaseSlider);
    Vue.component(BaseProgress.name, BaseProgress);
    Vue.component(BasePagination.name, BasePagination);
    Vue.component(Card.name, Card);
  }
};
