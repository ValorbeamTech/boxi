// Types for Blog Card 8
export interface Post {
  uid: string;
  cover: string;
  created_at: string;
  title: string;
  description: string;
  author: {
    name: string;
    avatar?: string;
  };
}

export const posts: Post[] = [
  {
    uid: "1",
    cover: "/images/800x600.png",
    created_at: "25 May, 2022",
    title: "Food: A Simple Definition",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Animi necessitatibus repellat voluptatibus?",
    author: {
      name: "John Doe",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "2",
    cover: "/images/800x600.png",
    created_at: "30 May, 2022",
    title: "Tailwind CSS Card Example",
    description:
      "Lorem ipsum dolor sit amet, consectetur. Lorem ipsum dolor on the sit.",
    author: {
      name: "Konnor Guzman",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "3",
    cover: "/images/800x600.png",
    created_at: "10 June, 2022",
    title: "What is PHP?",
    description:
      "How Did We Get Here? The History of Music Told Through Tweets",
    author: {
      name: "Travis Fuller",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "4",
    cover: "/images/800x600.png",
    created_at: "19 June, 2022",
    title: "Top Design Systems",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Architecto assumenda.",
    author: {
      name: "Alfredo Elliott",
      avatar: undefined,
    },
  },
  {
    uid: "5",
    cover: "/images/800x600.png",
    created_at: "26 June, 2022",
    title: "Top Design Systems",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Architecto assumenda.",
    author: {
      name: "Katrina West",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "6",
    cover: "/images/800x600.png",
    created_at: "12 June, 2022",
    title: "How Did We Get Here? The History of Music Told Through Tweets",
    description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
    author: {
      name: "Henry Curtis",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "7",
    cover: "/images/800x600.png",
    created_at: "30 June, 2022",
    title: "NodeJS Design Patterns",
    description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
    author: {
      name: "Samantha Shelton",
      avatar: undefined,
    },
  },
  {
    uid: "8",
    cover: "/images/800x600.png",
    created_at: "03 July, 2022",
    title: "313 Pattern and Color ideas",
    description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
    author: {
      name: "Konnor Guzman",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "9",
    cover: "/images/800x600.png",
    created_at: "08 July, 2022",
    title: "25 Surprising Facts About Chair",
    description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
    author: {
      name: "Selena Gomez",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "10",
    cover: "/images/800x600.png",
    created_at: "10 June, 2022",
    title: "What is PHP?",
    description:
      "How Did We Get Here? The History of Music Told Through Tweets",
    author: {
      name: "Travis Fuller",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "11",
    cover: "/images/800x600.png",
    created_at: "11 September, 2022",
    title: "Food: A Simple Definition",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Animi necessitatibus repellat voluptatibus?",
    author: {
      name: "John Doe",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "12",
    cover: "/images/800x600.png",
    created_at: "25 May, 2022",
    title: "Tailwind CSS Card Example",
    description:
      "Lorem ipsum dolor sit amet, consectetur. Lorem ipsum dolor on the sit.",
    author: {
      name: "Konnor Guzman",
      avatar: "/images/200x200.png",
    },
  },
];
