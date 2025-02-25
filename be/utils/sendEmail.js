const nodemailer = require('nodemailer');

const sendEmail = async (to, subject, text) => {
    console.log(process.env.EMAIL_USER, process.env.EMAIL_PASS);        
    console.log(process.env.NODE_MAILER_EMAIL, process.env.NODE_MAILER_PASS);
  const transporter = nodemailer.createTransport({
    service: 'Gmail',  
    auth: {
      user: process.env.EMAIL_USER,
      pass: process.env.EMAIL_PASS
    }
  });

  const mailOptions = {
    from: process.env.EMAIL_USER,
    to,
    subject,
    text
  };

  await transporter.sendMail(mailOptions);
};

module.exports = sendEmail;
