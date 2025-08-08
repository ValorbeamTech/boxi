export const colors = {
  primary: "#6366f1",
  secondary: "#F000B9",
  success: "#10B981",
  warning: "#FF9800",
  error: "#FF5724",
};

export type ColorVariant = keyof typeof colors;

export interface User {
  uid: string;
  cover: string;
  avatar?: string | null;
  color: ColorVariant;
  socialLinks: {
    twitter: string;
    instagram: string;
    facebook: string;
  };
  name: string;
  location: string;
  chartData: number[];
  postsCount: number;
}

export const users: User[] = [
  {
    uid: "1",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "primary",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Konnor Guzman",
    location: "USA, Washington DC",
    chartData: [48, 100, 70, 92],
    postsCount: 24,
  },
  {
    uid: "2",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "secondary",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Travis Fuller",
    location: "UK, London",
    chartData: [54, 77, 43, 69, 12],
    postsCount: 11,
  },
  {
    uid: "3",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "success",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Alfredo Elliott",
    location: "Australia, Sydney",
    chartData: [0, 100, 0],
    postsCount: 171,
  },
  {
    uid: "4",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "error",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Derrick Simmons",
    location: "Austria, Vienna",
    chartData: [0, 20, 10, 30, 20, 50],
    postsCount: 67,
  },
  {
    uid: "5",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "warning",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Raul Bradley",
    location: "Germany, Berlin",
    chartData: [33, 77, 55, 102, 12],
    postsCount: 36,
  },
  {
    uid: "6",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "primary",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Katrina West",
    location: "Sweden, Stockholm",
    chartData: [48, 100, 70, 92],
    postsCount: 52,
  },
  {
    uid: "7",
    cover: "/images/800x600.png",
    avatar: null,
    color: "secondary",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Samantha Shelton",
    location: "Switzerland, Zurich",
    chartData: [54, 77, 43, 69, 12],
    postsCount: 96,
  },
  {
    uid: "8",
    cover: "/images/800x600.png",
    avatar: "/images/200x200.png",
    color: "success",
    socialLinks: {
      twitter: "https://twitter.com",
      instagram: "https://instagram.com",
      facebook: "https://facebook.com",
    },
    name: "Henry Curtis",
    location: "Australia, Sydney",
    chartData: [0, 100, 0],
    postsCount: 125,
  },
];
