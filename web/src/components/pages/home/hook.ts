import { Space } from "@/types/interface";

export default () => {
  // temporary data
  const spaces: Space[] = [
    {
      id: 1,
      name: "Langara",
      href: "#",
      imageSrc:
        "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
      imageAlt: "Front of men's Basic Tee in black.",
      price: "$35",
      color: "Black",
      lat: 49.2244,
      lng: -123.1089,
    },
    {
      id: 2,
      name: "Whatever",
      href: "#",
      imageSrc:
        "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
      imageAlt: "Front of men's Basic Tee in black.",
      price: "$35",
      color: "Black",
      lat: 50.381832,
      lng: -120.623177,
    },
    {
      id: 3,
      name: "Los Angeles",
      href: "#",
      imageSrc:
        "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
      imageAlt: "Front of men's Basic Tee in black.",
      price: "$35",
      color: "Black",
      lat: 34.052235,
      lng: -118.243683,
    },
    {
      id: 4,
      name: "Home",
      href: "#",
      imageSrc:
        "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
      imageAlt: "Front of men's Basic Tee in black.",
      price: "$35",
      color: "Black",
      lat: 35.52389,
      lng: 139.69294,
    },
  ];
  return { spaces };
};
