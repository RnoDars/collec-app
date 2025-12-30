import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Collec-App - Gestion de collections",
  description: "Application moderne de gestion de collections",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="fr">
      <body>
        {children}
      </body>
    </html>
  );
}
