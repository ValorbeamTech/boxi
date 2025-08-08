export interface User {
  uid: string;
  name: string;
  avatar?: string | null;
  position: string;
}

export const users: User[] = [
  {
    uid: "1",
    name: "Konnor Guzman",
    avatar: "/images/200x200.png",
    position: "Senior Developer",
  },
  {
    uid: "2",
    name: "Travis Fuller",
    avatar: "/images/200x200.png",
    position: "Web Developer",
  },
  {
    uid: "3",
    name: "Alfredo Elliott",
    avatar: "/images/200x200.png",
    position: "UI/UX designer",
  },
  {
    uid: "4",
    name: "Derrick Simmons",
    avatar: null,
    position: "React Developer",
  },
  {
    uid: "5",
    name: "Katrina West",
    avatar: "/images/200x200.png",
    position: "Android Developer",
  },
  {
    uid: "6",
    name: "Henry Curtis",
    avatar: "/images/200x200.png",
    position: "Full Stack Developer",
  },
  {
    uid: "7",
    name: "Raul Bradley",
    avatar: "/images/200x200.png",
    position: "Laravel Developer",
  },
  {
    uid: "8",
    name: "Samantha Shelton",
    avatar: null,
    position: "Backend Developer",
  },
  {
    uid: "9",
    name: "Corey Evans",
    avatar: "/images/200x200.png",
    position: "Frontend Developer",
  },
  {
    uid: "10",
    name: "Lance Tucker",
    avatar: null,
    position: "NodeJS Developer",
  },
  {
    uid: "11",
    name: "Anthony Jensen",
    avatar: "/images/200x200.png",
    position: "UI/UX Designer",
  },
  {
    uid: "12",
    name: "Anthony Jensen",
    avatar: "/images/200x200.png",
    position: "Backend Developer",
  },
];
