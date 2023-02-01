import useTemplate from "./hook";
import Template from "./template";

const HomePageComponent = () => {
  const logic = useTemplate();
  return <Template {...logic} />;
};

HomePageComponent.displayName = "HomePageComponent";

export default HomePageComponent;
