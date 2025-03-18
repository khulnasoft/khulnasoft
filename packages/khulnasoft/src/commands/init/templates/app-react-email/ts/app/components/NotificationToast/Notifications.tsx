"use client";

import { Khulnasoft } from "@khulnasoft/js";
import { useEffect, useState, useMemo } from "react";
import { Inbox } from "@khulnasoft/nextjs";
import { useRouter } from "next/navigation";
import styles from "./Notifications.module.css"; // You'll need to create this

const NotificationToast = () => {
  const khulnasoft = useMemo(
    () =>
      new Khulnasoft({
        subscriberId: process.env.NEXT_PUBLIC_KHULNASOFT_SUBSCRIBER_ID || "",
        applicationIdentifier:
          process.env.NEXT_PUBLIC_KHULNASOFT_APPLICATION_IDENTIFIER || "",
      }),
    [],
  );

  const [showToast, setShowToast] = useState(false);

  useEffect(() => {
    const listener = ({ result: notification }: { result: any }) => {
      console.log("Received notification:", notification);
      setShowToast(true);

      setTimeout(() => {
        setShowToast(false);
      }, 2500);
    };

    console.log("Setting up Khulnasoft notification listener");
    khulnasoft.on("notifications.notification_received", listener);

    return () => {
      khulnasoft.off("notifications.notification_received", listener);
    };
  }, [khulnasoft]);

  if (!showToast) return null;

  return (
    <div className={styles.toast}>
      <div className={styles.toastContent}>New In-App Notification</div>
    </div>
  );
};

export default NotificationToast;

const khulnasoftConfig = {
  applicationIdentifier:
    process.env.NEXT_PUBLIC_KHULNASOFT_APPLICATION_IDENTIFIER || "",
  subscriberId: process.env.NEXT_PUBLIC_KHULNASOFT_SUBSCRIBER_ID || "",
  appearance: {
    elements: {
      bellContainer: {
        width: "30px",
        height: "30px",
      },
      bellIcon: {
        width: "30px",
        height: "30px",
      },
    },
  },
};

export function KhulnasoftInbox() {
  const router = useRouter();

  return (
    <Inbox {...khulnasoftConfig} routerPush={(path: string) => router.push(path)} />
  );
}
