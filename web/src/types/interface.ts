export interface Space {
  id: number;
  imageSrc: string;
  imageAlt: string;
  href: string;
  name: string;
  color: string;
  price: string;
  lat: number;
  lng: number;
}

export interface SignUpParams {
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
}

export interface SignInParams {
  emailAddress: string;
  password: string;
}

export interface User {
  email: string;
  firstName: string;
  lastName: string;
  imageUrl?: string;
}
