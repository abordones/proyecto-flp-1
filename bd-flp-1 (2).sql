-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1:3306
-- Tiempo de generación: 20-11-2022 a las 19:41:04
-- Versión del servidor: 5.7.36
-- Versión de PHP: 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `bd-flp-1`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `answers`
--

CREATE TABLE IF NOT EXISTS `answers` (
  `ID_A` int(11) NOT NULL AUTO_INCREMENT,
  `ID_Q` int(11) NOT NULL,
  `POINT_A` int(11) DEFAULT NULL,
  `OBSERVATION_A` varchar(200) DEFAULT NULL,
  `ACTIVE_A` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`ID_A`),
  UNIQUE KEY `ANSWERS_PK` (`ID_A`),
  KEY `TENER_FK` (`ID_Q`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `patients`
--

CREATE TABLE IF NOT EXISTS `patients` (
  `ID_P` int(11) NOT NULL AUTO_INCREMENT,
  `RUN_P` int(11) DEFAULT NULL,
  `DV_P` char(1) DEFAULT NULL,
  `NAME_P` char(100) DEFAULT NULL,
  `FATHERNAME_P` char(50) DEFAULT NULL,
  `MOTHERNAME_P` char(50) DEFAULT NULL,
  `BIRTHDAY_P` date DEFAULT NULL,
  `PHONE_P` int(11) DEFAULT NULL,
  `EMAIL_P` varchar(70) DEFAULT NULL,
  `OBSERVATION_P` varchar(500) DEFAULT NULL,
  `ACTIVE_P` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`ID_P`),
  UNIQUE KEY `PATIENTS_PK` (`ID_P`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `polls`
--

CREATE TABLE IF NOT EXISTS `polls` (
  `ID_PO` int(11) NOT NULL AUTO_INCREMENT,
  `ID_U` int(11) NOT NULL,
  `ID_T` int(11) NOT NULL,
  PRIMARY KEY (`ID_PO`),
  UNIQUE KEY `POLLS_PK` (`ID_PO`),
  KEY `REALIZAR_FK` (`ID_U`),
  KEY `ESTAR_FK` (`ID_T`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `questions`
--

CREATE TABLE IF NOT EXISTS `questions` (
  `ID_Q` int(11) NOT NULL AUTO_INCREMENT,
  `ID_T` int(11) NOT NULL,
  `QUESTION_Q` varchar(150) DEFAULT NULL,
  `DESCRIPTION_Q` varchar(200) DEFAULT NULL,
  `ACTIVE_Q` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`ID_Q`),
  UNIQUE KEY `QUESTIONS_PK` (`ID_Q`),
  KEY `PREGUNTAR_FK` (`ID_T`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `questions`
--

INSERT INTO `questions` (`ID_Q`, `ID_T`, `QUESTION_Q`, `DESCRIPTION_Q`, `ACTIVE_Q`) VALUES
(1, 1, 'Tristeza\r\n', '-\r\n', 1),
(2, 1, 'Pesimismo\r\n', '-\r\n', 1),
(3, 1, 'Fracaso\r\n', '-\r\n', 1),
(4, 1, 'Pérdida de Placer\r\n', '-\r\n', 1),
(5, 1, 'Sentimientos de Culpa\r\n', '-\r\n', 1),
(6, 1, 'Sentimientos de Castigo\r\n', '-\r\n', 1),
(7, 1, 'Disconformidad con uno mismo\r\n', '-\r\n', 1),
(8, 1, 'Autocrítica\r\n', '-\r\n', 1),
(9, 1, 'Pensamientos o Deseos Suicidas\r\n', '-\r\n', 1),
(10, 1, 'Llanto\r\n', '-\r\n', 1),
(11, 1, 'Agitación\r\n', '-\r\n', 1),
(12, 1, 'Pérdida de Interés\r\n', '-\r\n', 1),
(13, 1, 'Indecisión\r\n', '-\r\n', 1),
(14, 1, 'Desvalorización\r\n', '-\r\n', 1),
(15, 1, 'Pérdida de Energía\r\n', '-\r\n', 1),
(16, 1, 'Cambios en los Habitos de Sueño\r\n', '-\r\n', 1),
(17, 1, 'Irritabilidad\r\n', '-\r\n', 1),
(18, 1, 'Cambios en el Apetito\r\n', '-\r\n', 1),
(19, 1, 'Dificultad de Concentración\r\n', '-\r\n', 1),
(20, 1, 'Cansancio o Fatiga\r\n', '-\r\n', 1),
(21, 1, 'Pérdida de Interés en el Sexo\r\n', '-\r\n', 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `sessions`
--

CREATE TABLE IF NOT EXISTS `sessions` (
  `ID_P` int(11) NOT NULL,
  `ID_PO` int(11) NOT NULL,
  `DATE_S` date DEFAULT NULL,
  PRIMARY KEY (`ID_P`,`ID_PO`),
  UNIQUE KEY `SESSION_PK` (`ID_P`,`ID_PO`),
  KEY `SESSION2_FK` (`ID_PO`),
  KEY `SESSION_FK` (`ID_P`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tests`
--

CREATE TABLE IF NOT EXISTS `tests` (
  `ID_T` int(11) NOT NULL AUTO_INCREMENT,
  `NAME_T` char(100) DEFAULT NULL,
  `CUTPOINT_T` int(11) DEFAULT NULL,
  `MATCHPOINT_T` int(11) DEFAULT NULL,
  `DESCRIPTION_T` varchar(500) DEFAULT NULL,
  `ACTIVE_T` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`ID_T`),
  UNIQUE KEY `TESTS_PK` (`ID_T`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `tests`
--

INSERT INTO `tests` (`ID_T`, `NAME_T`, `CUTPOINT_T`, `MATCHPOINT_T`, `DESCRIPTION_T`, `ACTIVE_T`) VALUES
(1, 'Inventario de Beck\r\n', 5, 63, 'Cuestionario por defecto del Test de Beck sobre la depresión\r\n', 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE IF NOT EXISTS `users` (
  `ID_U` int(11) NOT NULL AUTO_INCREMENT,
  `RUN_U` int(11) DEFAULT NULL,
  `DV_U` char(1) DEFAULT NULL,
  `NAME_U` char(100) DEFAULT NULL,
  `FATHERNAME_U` char(50) DEFAULT NULL,
  `MOTHERNAME_U` char(50) DEFAULT NULL,
  `BIRTHDAY_U` date DEFAULT NULL,
  `EMAIL_U` varchar(70) DEFAULT NULL,
  `PASSWORD_U` varchar(33) DEFAULT NULL,
  `ACTIVE_U` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`ID_U`),
  UNIQUE KEY `USERS_PK` (`ID_U`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `answers`
--
ALTER TABLE `answers`
  ADD CONSTRAINT `FK_ANSWERS_TENER_QUESTION` FOREIGN KEY (`ID_Q`) REFERENCES `questions` (`ID_Q`);

--
-- Filtros para la tabla `polls`
--
ALTER TABLE `polls`
  ADD CONSTRAINT `FK_POLLS_ESTAR_TESTS` FOREIGN KEY (`ID_T`) REFERENCES `tests` (`ID_T`),
  ADD CONSTRAINT `FK_POLLS_REALIZAR_USERS` FOREIGN KEY (`ID_U`) REFERENCES `users` (`ID_U`);

--
-- Filtros para la tabla `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `FK_QUESTION_PREGUNTAR_TESTS` FOREIGN KEY (`ID_T`) REFERENCES `tests` (`ID_T`);

--
-- Filtros para la tabla `sessions`
--
ALTER TABLE `sessions`
  ADD CONSTRAINT `FK_SESSIONS_SESSIONS2_POLLS` FOREIGN KEY (`ID_PO`) REFERENCES `polls` (`ID_PO`),
  ADD CONSTRAINT `FK_SESSIONS_SESSIONS_PATIENTS` FOREIGN KEY (`ID_P`) REFERENCES `patients` (`ID_P`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
