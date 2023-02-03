import { Space } from "@/types/interface";

export default () => {
  const spaces: Space[] = [];
  const space: Space = {
    id: 1,
    name: "Basic Tee",
    href: "#",
    imageSrc:
      "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
    imageAlt: "Front of men's Basic Tee in black.",
    price: "$35",
    color: "Black",
  };
  // temporary data
  const generateMockSpaces = () => {
    for (let i = 0; i < 3; i++) {
      spaces.push(space);
    }
  };
  generateMockSpaces();
  return { spaces };
};
