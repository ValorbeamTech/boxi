export interface User {
  uid: string;
  name: string;
  avatar?: string | null;
  position: string;
  team: { uid: string; name: string; avatar?: string | null }[];
}

export const users: User[] = [
  {
    uid: "1",
    name: "Konnor Guzman",
    avatar: "/images/200x200.png",
    position: "Senior Developer",
    team: [
      {
        uid: "1",
        name: "Travis Fuller",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "John Doe",
        avatar: null,
      },
      {
        uid: "3",
        name: "Caleb Nolan",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "2",
    name: "Travis Fuller",
    avatar: "/images/200x200.png",
    position: "Web Developer",
    team: [
      {
        uid: "1",
        name: "Julie Morgan",
        avatar: null,
      },
      {
        uid: "2",
        name: "John Doe",
        avatar: "/images/200x200.png",
      },
      {
        uid: "3",
        name: "Alfredo Elliott",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "3",
    name: "Alfredo Elliott",
    avatar: "/images/200x200.png",
    position: "UI/UX designer",
    team: [
      {
        uid: "1",
        name: "Megan Fox",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Selena Gomez",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "4",
    name: "Derrick Simmons",
    avatar: null,
    position: "React Developer",
    team: [
      {
        uid: "1",
        name: "Leonardo DiCaprio",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Lynda Headey",
        avatar: null,
      },
      {
        uid: "3",
        name: "Majid Yahyaei",
        avatar: null,
      },
    ],
  },
  {
    uid: "5",
    name: "Katrina West",
    avatar: "/images/200x200.png",
    position: "Android Developer",
    team: [
      {
        uid: "1",
        name: "Barak Obama",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Donald Trump",
        avatar: "/images/200x200.png",
      },
      {
        uid: "3",
        name: "Selena Gomez",
        avatar: null,
      },
    ],
  },
  {
    uid: "6",
    name: "Henry Curtis",
    avatar: "/images/200x200.png",
    position: "Full Stack Developer",
    team: [
      {
        uid: "1",
        name: "Adam Sandler",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Emilia Clarke",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "7",
    name: "Raul Bradley",
    avatar: "/images/200x200.png",
    position: "Laravel Developer",
    team: [
      {
        uid: "1",
        name: "Travis Fuller",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "John Doe",
        avatar: null,
      },
      {
        uid: "3",
        name: "Caleb Nolan",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "8",
    name: "Samantha Shelton",
    avatar: null,
    position: "Backend Developer",
    team: [
      {
        uid: "1",
        name: "Julie Morgan",
        avatar: null,
      },
      {
        uid: "2",
        name: "John Doe",
        avatar: "/images/200x200.png",
      },
      {
        uid: "3",
        name: "Alfredo Elliott",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "9",
    name: "Corey Evans",
    avatar: "/images/200x200.png",
    position: "Frontend Developer",
    team: [
      {
        uid: "1",
        name: "Megan Fox",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Selena Gomez",
        avatar: "/images/200x200.png",
      },
    ],
  },
  {
    uid: "10",
    name: "Lance Tucker",
    avatar: null,
    position: "NodeJS Developer",
    team: [
      {
        uid: "1",
        name: "Leonardo DiCaprio",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Lynda Headey",
        avatar: null,
      },
      {
        uid: "3",
        name: "Majid Yahyaei",
        avatar: null,
      },
    ],
  },
  {
    uid: "11",
    name: "Anthony Jensen",
    avatar: "/images/200x200.png",
    position: "UI/UX Designer",
    team: [
      {
        uid: "1",
        name: "Barak Obama",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Donald Trump",
        avatar: "/images/200x200.png",
      },
      {
        uid: "3",
        name: "Selena Gomez",
        avatar: null,
      },
    ],
  },
  {
    uid: "12",
    name: "Anthony Jensen",
    avatar: "/images/200x200.png",
    position: "Backend Developer",
    team: [
      {
        uid: "1",
        name: "Adam Sandler",
        avatar: "/images/200x200.png",
      },
      {
        uid: "2",
        name: "Emilia Clarke",
        avatar: "/images/200x200.png",
      },
    ],
  },
];
