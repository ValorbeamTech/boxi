export interface Post {
  uid: string;
  cover: string;
  title: string;
  readTime: string;
  author: {
    name: string;
    avatar?: string | null;
  };
}

export const posts: Post[] = [
  {
    uid: "1",
    cover: "/images/800x600.png",
    title: "What is Tailwind CSS?",
    readTime: "2 min read",
    author: {
      name: "John D.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "2",
    cover: "/images/800x600.png",
    title: "Tailwind CSS Card Example",
    readTime: "5 min read",
    author: {
      name: "Travis F.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "3",
    cover: "/images/800x600.png",
    title: "10 Tips for Making a Good Camera Even Better",
    readTime: "4 min read",
    author: {
      name: "Alfredo E.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "4",
    cover: "/images/800x600.png",
    title: "Top Design Systems",
    readTime: "1 min read",
    author: {
      name: "Katrina W.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "5",
    cover: "/images/800x600.png",
    title: "NodeJS Design Patterns",
    readTime: "6 min read",
    author: {
      name: "Henry C.",
      avatar: null,
    },
  },
  {
    uid: "6",
    cover: "/images/800x600.png",
    title: "313 Pattern and Color ideas",
    readTime: "3 min read",
    author: {
      name: "Derrick S.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "7",
    cover: "/images/800x600.png",
    title: "10 Tips for Making a Good Camera Even Better",
    readTime: "1 min read",
    author: {
      name: "Raul B.",
      avatar: null,
    },
  },
  {
    uid: "8",
    cover: "/images/800x600.png",
    title: "The 12 Worst Types Business Accounts You Follow on Twitter",
    readTime: "7 min read",
    author: {
      name: "Samantha S.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "9",
    cover: "/images/800x600.png",
    title: "25 Surprising Facts About Chair",
    readTime: "2 min read",
    author: {
      name: "Travis F.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "10",
    cover: "/images/800x600.png",
    title: "What is Tailwind CSS?",
    readTime: "2 min read",
    author: {
      name: "John D.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "11",
    cover: "/images/800x600.png",
    title: "Tailwind CSS Card Example",
    readTime: "5 min read",
    author: {
      name: "Travis F.",
      avatar: "/images/200x200.png",
    },
  },
  {
    uid: "12",
    cover: "/images/800x600.png",
    title: "10 Tips for Making a Good Camera Even Better",
    readTime: "4 min read",
    author: {
      name: "Alfredo E.",
      avatar: "/images/200x200.png",
    },
  },
];
